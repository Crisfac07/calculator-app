import { useState } from 'react'
import type { CalculatorOperation, CalculatorToken } from '../../types/calculator'
import { evaluateExpression } from '../../services/expressionEvaluator'
import CalculatorButton from './CalculatorButton'
import CalculatorDisplay from './CalculatorDisplay'

function formatToken(token: CalculatorToken): string {
  if (token.type === 'number') {
    return token.value
  }

  const operationSymbols: Record<CalculatorOperation, string> = {
    add: '+',
    subtract: '−',
    multiply: '×',
    divide: '÷',
    power: 'xʸ',
    sqrt: '√',
  }

  return operationSymbols[token.value]
}

function Calculator() {

const [currentInput, setCurrentInput] = useState('')
const [tokens, setTokens] = useState<CalculatorToken[]>([])

const handleNumberClick = (number: string) => {
    setCurrentInput((previous) => previous + number)
  }
const handleDecimalClick = () => {
  if (currentInput.includes('.')) {
    return
  }

  setCurrentInput((previous) => {
    if (previous === '') {
      return '0.'
    }

    return previous + '.'
  })
}
const handleOperationClick = (operation: CalculatorOperation) => {
  if (currentInput === '') {
    setTokens((previous) => {
      if (previous.length === 0) {
        return previous
      }

      const lastToken = previous[previous.length - 1]

      if (lastToken.type === 'operation') {
        return [
          ...previous.slice(0, -1),
          {
            type: 'operation',
            value: operation,
          },
        ]
      }

      return [
        ...previous,
        {
          type: 'operation',
          value: operation,
        },
      ]
    })

    return
  }

  setTokens((previous) => [
    ...previous,
    {
      type: 'number',
      value: currentInput,
    },
    {
      type: 'operation',
      value: operation,
    },
  ])

  setCurrentInput('')
}

const handleSquareRootClick = () => {
  if (currentInput === '') {
    setTokens((previous) => {
      const lastToken = previous[previous.length - 1]

      if (lastToken?.type === 'operation') {
        return previous
      }

      return [
        ...previous,
        {
          type: 'operation',
          value: 'sqrt',
        },
      ]
    })

    return
  }

  setTokens((previous) => [
    ...previous,
    {
      type: 'operation',
      value: 'sqrt',
    },
    {
      type: 'number',
      value: currentInput,
    },
  ])

  setCurrentInput('')
}

const handleCalculate = async () => {
  if (currentInput === '') {
    return
  }

  const expressionTokens: CalculatorToken[] = [
    ...tokens,
    {
      type: 'number',
      value: currentInput,
    },
  ]

  try {
    const result = await evaluateExpression(expressionTokens)

    setCurrentInput(String(result))
    setTokens([])
  } catch {
    setCurrentInput('Error')
    setTokens([])
  }
}

const handleClear = () => {
  setCurrentInput('')
  setTokens([])
}

  const numberButtons = [
    '7', '8', '9',
    '4', '5', '6',
    '1', '2', '3',
    '0',
  ]

  return (
    <section className="card shadow-sm">
      <div className="card-body">
        <h2 className="h4 mb-4">Calculator</h2>

        <CalculatorDisplay
            value={`${tokens.map(formatToken).join(' ')} ${currentInput}`.trim() || '0'}
        />

        {/* Operation buttons */}
        <div className="row g-2">
          <div className="col-4 col-sm-2">
            <CalculatorButton
              label="+"
              onClick={() => handleOperationClick('add')}
            />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton
              label="−"
              onClick={() => handleOperationClick('subtract')}
            />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton
              label="×"
              onClick={() => handleOperationClick('multiply')}
            />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton
              label="÷"
              onClick={() => handleOperationClick('divide')}
            />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton
              label="xʸ"
              onClick={() => handleOperationClick('power')}
            />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton
              label="√"
              onClick={handleSquareRootClick}
            />
          </div>
        </div>

        {/* Number buttons */}
        <div className="row g-2 mt-1">
          {numberButtons.map((number) => (
            <div className="col-4" key={number}>
              <CalculatorButton
                label={number}
                variant="light"
                onClick={() => handleNumberClick(number)}
              />
            </div>
          ))}
        </div>

        <div className="row g-2 mt-1">
          <div className="col-4">
            <CalculatorButton
              label="C"
              variant="danger"
              onClick={handleClear}
            />
          </div>

          <div className="col-4">
            <CalculatorButton
              label="."
              variant="light"
              onClick={handleDecimalClick}
            />
          </div>

          <div className="col-4">
            <CalculatorButton
              label="="
              variant="success"
              onClick={handleCalculate}
            />
          </div>
        </div>
      </div>
    </section>
  )
}

export default Calculator