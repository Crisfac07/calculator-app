export type CalculatorOperation =
  | 'add'
  | 'subtract'
  | 'multiply'
  | 'divide'
  | 'power'
  | 'sqrt'

export interface CalculatorRequest {
  operation: CalculatorOperation
  operands: number[]
}

export interface CalculatorResponse {
  result: number
}

export interface CalculatorError {
  code: string
  message: string
}

export interface CalculatorErrorResponse {
  error: CalculatorError
}