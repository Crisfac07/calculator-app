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

function findHighestPrecedenceOperation(
  tokens: CalculatorToken[],
): number {
  let highestPrecedence = -1
  let operationIndex = -1

  for (let index = 1; index < tokens.length; index += 2) {
    const token = tokens[index]

    if (token.type !== 'operation') {
      throw new Error('invalid expression')
    }

    const precedence = getPrecedence(token.value)

    if (precedence > highestPrecedence) {
      highestPrecedence = precedence
      operationIndex = index
    }
  }

  return operationIndex
}

async function resolveOperationAtIndex(
  tokens: CalculatorToken[],
  operationIndex: number,
): Promise<CalculatorToken[]> {
  const operationToken = tokens[operationIndex]
  const leftToken = tokens[operationIndex - 1]
  const rightToken = tokens[operationIndex + 1]

  if (
    operationToken.type !== 'operation' ||
    leftToken.type !== 'number' ||
    rightToken.type !== 'number'
  ) {
    throw new Error('invalid expression')
  }

  const result = await executeOperation(
    operationToken.value,
    [
      Number(leftToken.value),
      Number(rightToken.value),
    ],
  )

  return [
    ...tokens.slice(0, operationIndex - 1),
    {
      type: 'number',
      value: String(result),
    },
    ...tokens.slice(operationIndex + 2),
  ]
}

async function resolveSquareRoots(
  tokens: CalculatorToken[],
): Promise<CalculatorToken[]> {
  let currentTokens = tokens

  for (let index = 0; index < currentTokens.length; index += 1) {
    const token = currentTokens[index]

    if (
      token.type !== 'operation' ||
      token.value !== 'sqrt'
    ) {
      continue
    }

    const operandToken = currentTokens[index + 1]

    if (
      !operandToken ||
      operandToken.type !== 'number'
    ) {
      throw new Error('invalid expression')
    }

    const result = await executeOperation(
      'sqrt',
      [Number(operandToken.value)],
    )

    currentTokens = [
      ...currentTokens.slice(0, index),
      {
        type: 'number',
        value: String(result),
      },
      ...currentTokens.slice(index + 2),
    ]

    index -= 1
  }

  return currentTokens
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

async function resolveBinaryExpression(
  tokens: CalculatorToken[],
): Promise<number> {
  if (tokens.length < 3 || tokens.length % 2 === 0) {
    throw new Error('invalid expression')
  }

  const firstToken = tokens[0]

  if (firstToken.type !== 'number') {
    throw new Error('invalid expression')
  }

  let result = Number(firstToken.value)

  for (let index = 1; index < tokens.length; index += 2) {
    const operationToken = tokens[index]
    const operandToken = tokens[index + 1]

    if (
      operationToken.type !== 'operation' ||
      operandToken.type !== 'number'
    ) {
      throw new Error('invalid expression')
    }

    result = await executeOperation(
      operationToken.value,
      [result, Number(operandToken.value)],
    )
  }

  return result
}

async function evaluateExpression(
  tokens: CalculatorToken[],
): Promise<number> {
  const tokensWithoutSquareRoots = await resolveSquareRoots(tokens)

  const resolvedTokens = await resolveByPrecedence(
    tokensWithoutSquareRoots,
  )

  if (resolvedTokens.length !== 1) {
    throw new Error('invalid expression')
  }

  const resultToken = resolvedTokens[0]

  if (resultToken.type !== 'number') {
    throw new Error('invalid expression')
  }

  return Number(resultToken.value)
}

async function resolveByPrecedence(
  tokens: CalculatorToken[],
): Promise<CalculatorToken[]> {
  let currentTokens = tokens

  while (currentTokens.length > 1) {
    const operationIndex = findHighestPrecedenceOperation(currentTokens)

    if (operationIndex === -1) {
      throw new Error('invalid expression')
    }

    currentTokens = await resolveOperationAtIndex(
      currentTokens,
      operationIndex,
    )
  }

  return currentTokens
}

export {
  getPrecedence,
  executeOperation,
  evaluateExpression,
  resolveBinaryExpression,
  findHighestPrecedenceOperation,
  resolveOperationAtIndex,
  resolveByPrecedence,
  resolveSquareRoots
}