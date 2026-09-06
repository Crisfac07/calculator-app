import './App.css'
import Calculator from './components/Calculator/Calculator'

function App() {
  return (
    <main className="container py-5">
      <header className="mb-4">
        <h1 className="fw-bold">Calculator App</h1>
      </header>

      <div className="row g-4">
        <div className="col-12 col-lg-7">
          <Calculator />
        </div>

        <div className="col-12 col-lg-5">
          <section className="card shadow-sm h-100">
            <div className="card-body">
              <h2 className="h4">How to use</h2>

              <ol className="mt-3">
                <li>Enter a number.</li>
                <li>Select an operation.</li>
                <li>Enter the required number(s).</li>
                <li>Repeat the operation if needed.</li>
                <li>Press = to calculate.</li>
                <li>Press C to clear the current calculation.</li>
              </ol>
              <h3 className="h6 mt-4">Example</h3>
              <ul className="mb-0">
                <li>2 × 3 + 4 =</li>
              </ul>

              <h3 className="h6 mt-4">Supported operations</h3>

              <ul className="mb-0">
                <li>Addition</li>
                <li>Subtraction</li>
                <li>Multiplication</li>
                <li>Division</li>
                <li>Power</li>
                <li>Square root</li>
              </ul>
            </div>
          </section>
        </div>
      </div>
    </main>
  )
}

export default App