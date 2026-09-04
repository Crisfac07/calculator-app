import { calculate } from './calculator'
import type {
  CalculatorOperation,
  CalculatorToken,
} from '../types/calculator'

function getPrecedence(operation: CalculatorOperation): number {
  switch (operation) {
    case 'add':
    case 'subtract':
      return 1

    case 'multiply':
    case 'divide':
      return 2

    case 'power':
      return 3

    case 'sqrt':
      return 4
  }
}

async function executeOperation(
  operation: CalculatorOperation,
  operands: number[],
): Promise<number> {
  const response = await calculate({
    operation,
    operands,
  })

  return response.result
}

async function evaluateExpression(
  tokens: CalculatorToken[],
): Promise<number> {
  if (tokens.length !== 3) {
    throw new Error('invalid expression')
  }

  const [leftToken, operationToken, rightToken] = tokens

  if (
    leftToken.type !== 'number' ||
    operationToken.type !== 'operation' ||
    rightToken.type !== 'number'
  ) {
    throw new Error('invalid expression')
  }

  const left = Number(leftToken.value)
  const right = Number(rightToken.value)

  return executeOperation(
    operationToken.value,
    [left, right],
  )
}

export {
  getPrecedence,
  executeOperation,
  evaluateExpression,
}