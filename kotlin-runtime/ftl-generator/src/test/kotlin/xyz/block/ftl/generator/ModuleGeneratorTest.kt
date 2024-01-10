package xyz.block.ftl.generator

import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import xyz.block.ftl.v1.schema.*
import xyz.block.ftl.v1.schema.Array
import xyz.block.ftl.v1.schema.Float
import xyz.block.ftl.v1.schema.Int
import xyz.block.ftl.v1.schema.Map
import xyz.block.ftl.v1.schema.String
import kotlin.test.assertEquals

class ModuleGeneratorTest {
  private lateinit var generator: ModuleGenerator

  @BeforeEach
  fun setUp() {
    generator = ModuleGenerator()
  }

  @Test
  fun `should generate basic module`() {
    val file = generator.generateModule(Module(name = "test"))
    val expected = """// Code generated by FTL-Generator, do not edit.
//
package ftl.test

import xyz.block.ftl.Ignore

@Ignore
public class TestModule()
"""
    assertEquals(expected, file.toString())
  }

  @Test
  fun `should generate all Types`() {
    val decls = listOf(
      Decl(data_ = Data(comments = listOf("Request comments"), name = "TestRequest")),
      Decl(
        data_ = Data(
          comments = listOf("Response comments"), name = "TestResponse", fields = listOf(
            Field(name = "int", type = Type(int = Int())),
            Field(name = "float", type = Type(float = Float())),
            Field(name = "string", type = Type(string = String())),
            Field(name = "bytes", type = Type(bytes = Bytes())),
            Field(name = "bool", type = Type(bool = Bool())),
            Field(name = "time", type = Type(time = Time())),
            Field(name = "optional", type = Type(optional = Optional(type = Type(string = String())))),
            Field(name = "array", type = Type(array = Array(element = Type(string = String())))),
            Field(
              name = "nestedArray", type = Type(
                array = Array(element = Type(array = Array(element = Type(string = String()))))
              )
            ),
            Field(
              name = "dataRefArray", type = Type(
                array = Array(element = Type(dataRef = DataRef(name = "TestRequest", module = "test")))
              )
            ),
            Field(
              name = "map",
              type = Type(map = Map(key = Type(string = String()), value_ = Type(int = Int())))
            ),
            Field(
              name = "nestedMap", type = Type(
                map = Map(
                  key = Type(string = String()),
                  value_ = Type(map = Map(key = Type(string = String()), value_ = Type(int = Int())))
                )
              )
            ),
            Field(name = "dataRef", type = Type(dataRef = DataRef(name = "TestRequest"))),
            Field(name = "externalDataRef", type = Type(dataRef = DataRef(module = "other", name = "TestRequest"))),
          )
        )
      ),
    )
    val module = Module(name = "test", comments = listOf("Module comments"), decls = decls)

    val file = generator.generateModule(module)
    val expected = """// Code generated by FTL-Generator, do not edit.
// Module comments
package ftl.test

import java.time.OffsetDateTime
import kotlin.Boolean
import kotlin.ByteArray
import kotlin.Float
import kotlin.Long
import kotlin.String
import kotlin.Unit
import kotlin.collections.ArrayList
import kotlin.collections.Map
import xyz.block.ftl.Ignore

/**
 * Request comments
 */
public data class TestRequest(
  public val _empty: Unit = Unit,
)

/**
 * Response comments
 */
public data class TestResponse(
  public val int: Long,
  public val float: Float,
  public val string: String,
  public val bytes: ByteArray,
  public val bool: Boolean,
  public val time: OffsetDateTime,
  public val optional: String? = null,
  public val array: ArrayList<String>,
  public val nestedArray: ArrayList<ArrayList<String>>,
  public val dataRefArray: ArrayList<TestRequest>,
  public val map: Map<String, Long>,
  public val nestedMap: Map<String, Map<String, Long>>,
  public val dataRef: TestRequest,
  public val externalDataRef: ftl.other.TestRequest,
)

@Ignore
public class TestModule()
"""
    assertEquals(expected, file.toString())
  }

  @Test
  fun `should generate all Verbs`() {
    val decls = listOf(
      Decl(data_ = Data(comments = listOf("Request comments"), name = "TestRequest")),
      Decl(data_ = Data(comments = listOf("Response comments"), name = "TestResponse")),
      Decl(
        verb = Verb(
          name = "TestVerb",
          comments = listOf("TestVerb comments"),
          request = DataRef(name = "TestRequest"),
          response = DataRef(name = "TestResponse")
        )
      ),
      Decl(
        verb = Verb(
          name = "TestIngressVerb",
          comments = listOf("TestIngressVerb comments"),
          request = DataRef(name = "TestRequest"),
          response = DataRef(name = "TestResponse"),
          metadata = listOf(
            Metadata(
              ingress = MetadataIngress(
                type = "ftl",
                path = listOf(IngressPathComponent(ingressPathLiteral = IngressPathLiteral(text = "test"))),
                method = "GET"
              )
            ),
          )
        )
      ),
    )
    val module = Module(name = "test", comments = listOf("Module comments"), decls = decls)
    val file = generator.generateModule(module)
    val expected = """// Code generated by FTL-Generator, do not edit.
// Module comments
package ftl.test

import kotlin.Unit
import xyz.block.ftl.Context
import xyz.block.ftl.Ignore
import xyz.block.ftl.Ingress
import xyz.block.ftl.Method.GET
import xyz.block.ftl.Verb

/**
 * Request comments
 */
public data class TestRequest(
  public val _empty: Unit = Unit,
)

/**
 * Response comments
 */
public data class TestResponse(
  public val _empty: Unit = Unit,
)

@Ignore
public class TestModule() {
  /**
   * TestVerb comments
   */
  @Verb
  public fun TestVerb(context: Context, req: TestRequest): TestResponse = throw
      NotImplementedError("Verb stubs should not be called directly, instead use context.call(TestModule::TestVerb, ...)")

  /**
   * TestIngressVerb comments
   */
  @Verb
  @Ingress(
    GET,
    "/test",
  )
  public fun TestIngressVerb(context: Context, req: TestRequest): TestResponse = throw
      NotImplementedError("Verb stubs should not be called directly, instead use context.call(TestModule::TestIngressVerb, ...)")
}
"""
    assertEquals(expected, file.toString())
  }
}
