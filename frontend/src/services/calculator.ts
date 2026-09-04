import type {
  CalculatorErrorResponse,
  CalculatorRequest,
  CalculatorResponse,
} from '../types/calculator'

export async function calculate(
  request: CalculatorRequest,
): Promise<CalculatorResponse> {
  const response = await fetch('/api/v1/calculator', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(request),
  })

  if (!response.ok) {
    const error: CalculatorErrorResponse = await response.json()
    throw new Error(error.error.message)
  }

  return response.json()
}