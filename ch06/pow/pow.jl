"""
Compare three ways of computing a^n.

Run once to install the only non-standard dependency:
    import Pkg; Pkg.add("Plots")

Then run:
    julia power_experiment.jl

The actual arithmetic is performed modulo MOD.  This prevents the value of a^n
from becoming an enormous BigInt, so the timing mostly reflects the algorithms
rather than the cost of storing ever-larger integers.
"""

using Statistics
using DelimitedFiles
using Plots

const MOD = Int64(1_000_000_007)
const BASE = Int64(3)

# (1) O(n): multiplication in a loop
function pow_iterative(a::Int64, n::Int)
  value = Int64(1)
  multiplications = 0
  for _ in 1:n
    value = (value * a) % MOD
    multiplications += 1
  end
  return value, multiplications
end

# (2) O(n): the direct recursive definition a^n = a * a^(n-1)
function pow_recursive_naive(a::Int64, n::Int)
  n == 0 && return Int64(1), 0
  value, multiplications = pow_recursive_naive(a, n - 1)
  return (a * value) % MOD, multiplications + 1
end

# (3) O(log n): exponentiation by squaring
function pow_recursive_fast(a::Int64, n::Int)
  n == 0 && return Int64(1), 0
  n == 1 && return a % MOD, 0

  half, multiplications = pow_recursive_fast(a, n ÷ 2)
  square = (half * half) % MOD
  multiplications += 1             # the squaring: half * half

  if isodd(n)
    return (square * a) % MOD, multiplications + 1
  else
    return square, multiplications
  end
end

const METHODS = [
  ("iterative", pow_iterative),
  ("naive_recursive", pow_recursive_naive),
  ("fast_recursive", pow_recursive_fast),
]

"""Median wall-clock time in nanoseconds, after a warm-up compilation call."""
function median_time_ns(f, a::Int64, n::Int; repeats::Int=9)
  f(a, n) # compile before timing
  samples = Float64[]
  for _ in 1:repeats
    start = time_ns()
    result = f(a, n)
    elapsed = time_ns() - start
    # Keeps the computation observable; also checks every method agrees.
    result[1] == 0 && error("unexpected result")
    push!(samples, elapsed)
  end
  return median(samples)
end

function main()
  # Keep this below the stack limit of typical machines for the linear recursion.
  ns = [10, 30, 100, 300, 1_000, 3_000, 10_000]
  rows = Vector{Tuple{Int,String,Int,Float64}}()

  for n in ns
    answers = Int64[]
    for (name, f) in METHODS
      answer, count = f(BASE, n)
      push!(answers, answer)
      elapsed = median_time_ns(f, BASE, n)
      push!(rows, (n, name, count, elapsed))
      println(rpad(name, 18), " n=", lpad(n, 5),
        "  multiplications=", lpad(count, 5),
        "  median=", round(elapsed / 1e3; digits=2), " μs")
    end
    length(unique(answers)) == 1 || error("methods disagree at n=$n")
  end

  output_dir = @__DIR__
  csv_path = joinpath(output_dir, "power_experiment_results.csv")
  open(csv_path, "w") do io
    writedlm(io, ["n" "method" "multiplications" "median_time_ns"], ',')
    for row in rows
      writedlm(io, reshape(collect(row), 1, :), ',')
    end
  end

  labels = [method[1] for method in METHODS]
  count_plot = plot(xscale=:log10, yscale=:log10,
    xlabel="exponent n", ylabel="number of multiplications",
    title="Multiplications used to compute 3^n (mod $MOD)", legend=:topleft)
  time_plot = plot(xscale=:log10, yscale=:log10,
    xlabel="exponent n", ylabel="median execution time (ns)",
    title="Execution time: three exponentiation algorithms", legend=:topleft)

  for label in labels
    selected = filter(row -> row[2] == label, rows)
    x = [row[1] for row in selected]
    counts = [row[3] for row in selected]
    times = [row[4] for row in selected]
    plot!(count_plot, x, counts; marker=:circle, linewidth=2, label=label)
    plot!(time_plot, x, times; marker=:circle, linewidth=2, label=label)
  end

  savefig(count_plot, joinpath(output_dir, "multiplication_counts.png"))
  savefig(time_plot, joinpath(output_dir, "execution_times.png"))
  println("\nSaved CSV and two PNG graphs in: $output_dir")
end

main()
