import { beforeEach, describe, expect, it, vi } from 'vitest'

import {
    evaluateExpression,
  executeOperation,
  resolveByPrecedence,
} from './expressionEvaluator'

import { calculate } from './calculator'

vi.mock('./calculator', () => ({
  calculate: vi.fn(),
}))

beforeEach(() => {
  vi.mocked(calculate).mockReset()
})

describe('executeOperation', () => {
  it('should add two numbers', async () => {
    vi.mocked(calculate).mockResolvedValue({
      result: 5,
    })

    const result = await executeOperation('add', [2, 3])

    expect(result).toBe(5)
  })

  it('should subtract two numbers', async () => {
    vi.mocked(calculate).mockResolvedValue({
      result: 2,
    })

    const result = await executeOperation('subtract', [5, 3])

    expect(result).toBe(2)
  })

  it('should multiply two numbers', async () => {
    vi.mocked(calculate).mockResolvedValue({
      result: 15,
    })

    const result = await executeOperation('multiply', [5, 3])

    expect(result).toBe(15)
  })

  it('should divide two numbers', async () => {
    vi.mocked(calculate).mockResolvedValue({
      result: 2,
    })

    const result = await executeOperation('divide', [6, 3])

    expect(result).toBe(2)
  })

  it('should calculate a power', async () => {
    vi.mocked(calculate).mockResolvedValue({
      result: 8,
    })

    const result = await executeOperation('power', [2, 3])

    expect(result).toBe(8)
  })

  it('should calculate a square root', async () => {
    vi.mocked(calculate).mockResolvedValue({
      result: 12,
    })

    const result = await executeOperation('sqrt', [144])

    expect(result).toBe(12)
  })

  it('should propagate division by zero error', async ()=>{
    vi.mocked(calculate).mockRejectedValue(
        new Error('division by zero')
    )

    await expect(
        executeOperation('divide',[10,0]),
    ).rejects.toThrow('division by zero')
  })
})

describe('resolveByPrecedence', () => {
  it('should resolve multiplication before addition', async () => {
    vi.mocked(calculate)
      .mockResolvedValueOnce({
        result: 12,
      })
      .mockResolvedValueOnce({
        result: 14,
      })

    const tokens = [
      { type: 'number', value: '2' },
      { type: 'operation', value: 'add' },
      { type: 'number', value: '3' },
      { type: 'operation', value: 'multiply' },
      { type: 'number', value: '4' },
    ] as const

    const result = await resolveByPrecedence([...tokens])

    expect(result).toEqual([
      { type: 'number', value: '14' },
    ])

    expect(calculate).toHaveBeenNthCalledWith(1, {
      operation: 'multiply',
      operands: [3, 4],
    })

    expect(calculate).toHaveBeenNthCalledWith(2, {
      operation: 'add',
      operands: [2, 12],
    })
  })

  it('should resolve multiplication before addition when multiplication comes first', async () => {
    vi.mocked(calculate)
      .mockResolvedValueOnce({
        result: 6,
      })
      .mockResolvedValueOnce({
        result: 10,
      })

    const tokens = [
      { type: 'number', value: '2' },
      { type: 'operation', value: 'multiply' },
      { type: 'number', value: '3' },
      { type: 'operation', value: 'add' },
      { type: 'number', value: '4' },
    ] as const

    const result = await resolveByPrecedence([...tokens])

    expect(result).toEqual([
      { type: 'number', value: '10' },
    ])

    expect(calculate).toHaveBeenNthCalledWith(1, {
      operation: 'multiply',
      operands: [2, 3],
    })

    expect(calculate).toHaveBeenNthCalledWith(2, {
      operation: 'add',
      operands: [6, 4],
    })
  })
})

describe('evaluateExpression', () => {
  it('should evaluate square root and respect operator precedence', async () => {
    vi.mocked(calculate)
      .mockResolvedValueOnce({
        result: 12,
      })
      .mockResolvedValueOnce({
        result: 6,
      })
      .mockResolvedValueOnce({
        result: 18,
      })

    const tokens = [
      { type: 'operation', value: 'sqrt' },
      { type: 'number', value: '144' },
      { type: 'operation', value: 'add' },
      { type: 'number', value: '2' },
      { type: 'operation', value: 'multiply' },
      { type: 'number', value: '3' },
    ] as const

    const result = await evaluateExpression([...tokens])

    expect(result).toBe(18)

    expect(calculate).toHaveBeenNthCalledWith(1, {
      operation: 'sqrt',
      operands: [144],
    })

    expect(calculate).toHaveBeenNthCalledWith(2, {
      operation: 'multiply',
      operands: [2, 3],
    })

    expect(calculate).toHaveBeenNthCalledWith(3, {
      operation: 'add',
      operands: [12, 6],
    })
  })
})