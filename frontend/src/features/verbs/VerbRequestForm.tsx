import { useEffect, useState } from 'react'
import { Module, Verb } from '../../protos/xyz/block/ftl/v1/console/console_pb'
import { Ref } from '../../protos/xyz/block/ftl/v1/schema/schema_pb'
import { classNames } from '../../utils'
import { CodeEditor, InitialState } from '../../components/CodeEditor'
import { useClient } from '../../hooks/use-client'
import { VerbService } from '../../protos/xyz/block/ftl/v1/ftl_connect'
import { VerbFormInput } from './VerbFormInput'
import { createVerbRequest, defaultRequest, httpPopulatedRequestPath, isHttpIngress, requestPath, requestType, simpleJsonSchema } from './verb.utils'

export const VerbRequestForm = ({ module, verb }: { module?: Module; verb?: Verb }) => {
  const client = useClient(VerbService)
  const [activeTabId, setActiveTabId] = useState('body')
  const [initialEditorState, setInitialEditorText] = useState<InitialState>({ initialText: '' })
  const [editorText, setEditorText] = useState('')
  const [initialHeadersState, setInitialHeadersText] = useState<InitialState>({ initialText: '' })
  const [headersText, setHeadersText] = useState('')
  const [response, setResponse] = useState<string | null>(null)
  const [error, setError] = useState<string | null>(null)

  const editorTextKey = `${module?.name}-${verb?.verb?.name}-editor-text`

  useEffect(() => {
    if (verb) {
      const savedValue = localStorage.getItem(editorTextKey)
      let value: string
      if (savedValue != null && savedValue !== '') {
        value = savedValue
      } else {
        value = defaultRequest(verb)
      }

      const schemaString = JSON.stringify(simpleJsonSchema(verb))
      setInitialEditorText({ initialText: value, schema: schemaString })
      localStorage.setItem(editorTextKey, value)
      handleEditorTextChanged(value)

      const headerText = '{\n  "console": ["example"]\n}'
      setInitialHeadersText({ initialText: headerText })
      setHeadersText(headerText)
    }
  }, [verb, activeTabId])

  const handleEditorTextChanged = (text: string) => {
    setEditorText(text)
    localStorage.setItem(editorTextKey, text)
  }

  const handleTabClick = (e: React.MouseEvent<HTMLAnchorElement, MouseEvent>, id: string) => {
    e.preventDefault()
    setActiveTabId(id)
  }

  const tabs = [
    { id: 'body', name: 'Body' },
    { id: 'jsonschema', name: 'JSONSchema' }
  ]

  if (isHttpIngress(verb)) {
    tabs.push({ id: 'headers', name: 'Headers' })
  }

  const handleSubmit = async (path: string) => {
    setResponse(null)
    setError(null)

    try {
      const verbRef: Ref = {
        name: verb?.verb?.name,
        module: module?.name,
      } as Ref

      const requestBytes = createVerbRequest(path, verb, editorText, headersText)
      const response = await client.call({ verb: verbRef, body: requestBytes })
      if (response.response.case === 'body') {
        const jsonString = Buffer.from(response.response.value).toString('utf-8')

        setResponse(JSON.stringify(JSON.parse(jsonString), null, 2))
      } else if (response.response.case === 'error') {
        setError(response.response.value.message)
      }
    } catch (error) {
      console.error('There was an error with the request:', error)
      setError(String(error))
    }
  }

  return (
    <div className='p-4 flex flex-col h-full w-full'>
      <VerbFormInput
        requestType={requestType(verb)}
        initialPath={httpPopulatedRequestPath(module, verb)}
        requestPath={requestPath(module, verb)}
        readOnly={!isHttpIngress(verb)}
        onSubmit={handleSubmit}
      />
      <div>

        <div className='hidden sm:block pt-2'>
          <div className='border-b border-gray-200 dark:border-white/10'>
            <nav className='-mb-px flex space-x-6' aria-label='Tabs'>
              {tabs.map((tab) => (
                <a
                  key={tab.name}
                  className={classNames(
                    activeTabId === tab.id
                      ? 'border-indigo-500 text-indigo-600 dark:border-indigo-400 dark:text-indigo-400'
                      : 'border-transparent text-gray-500 hover:border-gray-300 hover:text-gray-700 dark:hover:border-gray-500 dark:text-gray-500 dark:hover:text-gray-300',
                    'whitespace-nowrap cursor-pointer border-b-2 py-2 px-1 text-sm font-medium'
                  )}
                  aria-current={activeTabId === tab.id ? 'page' : undefined}
                  onClick={(e) => handleTabClick(e, tab.id)}
                >
                  {tab.name}
                </a>
              ))}
            </nav>
          </div>
        </div>
      </div>
      <div className='flex flex-col justify-between h-full'>
        <div>
          {activeTabId === 'body' && (
            <CodeEditor key='body' initialState={initialEditorState} onTextChanged={handleEditorTextChanged} />
          )}
          {activeTabId === 'headers' && (
            <CodeEditor key='headers' initialState={initialHeadersState} onTextChanged={setHeadersText} />
          )}
          {activeTabId === 'jsonschema' && (
            <CodeEditor key='jsonschema' initialState={{ initialText: verb?.jsonRequestSchema ?? '', readonly: true }} />
          )}
        </div>

        <div className='flex flex-col'>
          {response && (
            <>
              <div className='border-b my-4 border-gray-200 dark:border-gray-600'></div>
              <h2 className='text-md pb-2'>Response</h2>
              <CodeEditor key='verb-response' initialState={{ initialText: response, readonly: true }} />
            </>
          )}
          {error && (
            <>
              <div className='border-b my-4 border-gray-200 dark:border-gray-600'></div>
              <h2 className='text-md pb-2'>Error</h2>
              <CodeEditor key='verb-error' initialState={{ initialText: error, readonly: true }} />
            </>
          )}
        </div>
      </div>
    </div>)
}