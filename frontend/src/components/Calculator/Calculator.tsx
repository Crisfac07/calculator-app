import CalculatorButton from "./CalculatorButton";
import CalculatorDisplay from "./CalculatorDisplay";

function Calculator() {
  const numberButtons = ["7", "8", "9", "4", "5", "6", "1", "2", "3", "0"];

  return (
    <section className="card shadow-sm">
      <div className="card-body">
        <h2 className="h4 mb-4">Calculator</h2>

        <CalculatorDisplay value="15" />

        {/* Operation buttons */}
        <div className="row g-2">
          <div className="col-4 col-sm-2">
            <CalculatorButton label="+" onClick={() => console.log("Add")} />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton
              label="−"
              onClick={() => console.log("Subtract")}
            />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton
              label="×"
              onClick={() => console.log("Multiply")}
            />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton label="÷" onClick={() => console.log("Divide")} />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton label="xʸ" onClick={() => console.log("Power")} />
          </div>

          <div className="col-4 col-sm-2">
            <CalculatorButton
              label="√"
              onClick={() => console.log("Square root")}
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
                onClick={() => console.log(number)}
              />
            </div>
          ))}
        </div>

        <div className="row g-2 mt-1">
          <div className="col-4">
            <CalculatorButton
              label="C"
              variant="danger"
              onClick={() => console.log("Clear")}
            />
          </div>

          <div className="col-4">
            <CalculatorButton
              label="."
              variant="light"
              onClick={() => console.log("Decimal")}
            />
          </div>

          <div className="col-4">
            <CalculatorButton
              label="="
              variant="success"
              onClick={() => console.log("Calculate")}
            />
          </div>
        </div>
      </div>
    </section>
  );
}

export default Calculator;
