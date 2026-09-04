interface CalculatorButtonProps {
  label: string
  onClick: () => void
  variant?: 'primary' | 'light' | 'success' | 'danger'
}

function CalculatorButton({
  label,
  onClick,
  variant = 'primary',
}: CalculatorButtonProps) {
  return (
    <button
      type="button"
      className={`btn btn-${variant} w-100`}
      onClick={onClick}
    >
      {label}
    </button>
  )
}

export default CalculatorButton