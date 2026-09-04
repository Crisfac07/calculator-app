interface CalculatorDisplayProps {
  value: string
}

function CalculatorDisplay({ value }: CalculatorDisplayProps) {
  return (
    <div className="border rounded p-3 mb-4 text-end">
      <span className="fs-2">{value}</span>
    </div>
  )
}

export default CalculatorDisplay