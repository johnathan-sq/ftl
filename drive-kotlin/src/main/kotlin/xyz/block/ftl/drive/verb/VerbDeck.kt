package xyz.block.ftl.drive.verb

import io.github.classgraph.ClassGraph
import xyz.block.ftl.Context
import xyz.block.ftl.Verb
import xyz.block.ftl.drive.Logging
import xyz.block.ftl.drive.verb.probe.ArgumentTracingProbe
import xyz.block.ftl.drive.verb.probe.TracingProbe
import java.util.concurrent.ConcurrentHashMap
import java.util.concurrent.CopyOnWriteArrayList
import kotlin.reflect.KClass
import kotlin.reflect.KFunction
import kotlin.reflect.KFunction1
import kotlin.reflect.jvm.kotlinFunction

class VerbDeck {
  private val logger = Logging.logger(VerbDeck::class)

  companion object {
    val instance = VerbDeck()

    fun init(module: String) {
      instance.init(module)
    }
  }

  data class VerbId(val qualifiedName: String)
  data class VerbSignature(val verbId: VerbId, val argumentType: KClass<*>)

  var module: String? = null
    private set
  private val verbs = ConcurrentHashMap<VerbId, VerbCassette<out Any>>()
  private val probes = CopyOnWriteArrayList<TracingProbe>()

  fun init(module: String) {
    this.module = module
    logger.info("Scanning for Verbs in ${module}...")
    ClassGraph()
      .enableAllInfo() // Scan classes, methods, fields, annotations
      .acceptPackages(module)
      .disableJarScanning()
      .scan().use { scanResult ->
        // Use the ScanResult within the try block, e.g.
        for (clazz in scanResult.getClassesWithMethodAnnotation(Verb::class.java)) {
          val kClass = clazz.loadClass().kotlin
          logger.info("    ${kClass.qualifiedName}")

          clazz.methodInfo
            .filter { info -> info.hasAnnotation(Verb::class.java) }
            .forEach { info ->
              logger.info("      @Verb ${info.name}()")
              val function = info.loadClassAndGetMethod().kotlinFunction!!

              val verbId = toId(function)
              @Suppress("UNCHECKED_CAST")
              verbs[verbId] = VerbCassette(verbId, kClass, function as KFunction1<Any, Any>)
            }
        }
      }

    // Now go through all the verbs and process their schema.
    logger.info("Registering schemas...")
    verbs.values.forEach { verb ->
      verb.readSchema()
    }

    probes.add(ArgumentTracingProbe())

    logger.info("Probes currently deployed: $probes")
  }

  fun fullyQualifiedName(id: VerbId): String = module!! + "." + id.qualifiedName

  fun list(): Set<VerbId> = verbs.keys

  fun lookup(name: String): VerbSignature? {
    val verbId = VerbId(name)

    return verbs[verbId]?.toDescriptor()
  }

  fun lookupFullyQualifiedName(name: String): VerbSignature? {
    val prefix = module!! + "."
    val verbId = VerbId(name.removePrefix(prefix))

    return if (name.startsWith(prefix)) {
      verbs[verbId]?.toDescriptor()
    } else null
  }

  fun dispatch(context: Context, verb: KFunction<*>, request: Any): Any {
    val verbId = toId(verb)
    return dispatch(Context.fromLocal(verbId, context), verbId, request)
  }

  fun dispatch(context: Context, verbId: VerbId, request: Any): Any {
    val cassette = verbs[verbId]!!

    // apply any probes that are in effect prior to dispatch
    probes.forEach { probe -> probe.probe(cassette, request, context.trace) }

    logger.debug("Local dispatch of ${verbId} [trace: ${context.trace}]")
    return cassette.invokeVerbInternal(context, request)
  }

  private fun toId(verb: KFunction<*>) = VerbId(verb.name)
}
