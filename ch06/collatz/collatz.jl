using Plots

collatz(a) = iseven(a) ? a ÷ 2 : 3a + 1

function collatz_stats(n)
  a = n
  term = 1
  val_max = a

  while a != 1
    a = collatz(a)
    term += 1
    val_max = max(a, val_max)
  end

  return term, val_max
end

function main()
  n_max = 200

  terms = zeros(Int, n_max)
  vals_max = zeros(Int, n_max)

  for i in 1:n_max
    terms[i], vals_max[i] = collatz_stats(i)
  end

  p1 = bar(1:n_max, terms, legend=false, title="Terms")
  p2 = bar(1:n_max, vals_max, legend=false, title="Max Values")

  p = plot(p1, p2, layout=(1, 2), size=(2000, 1000))
  savefig(p, "collatz_results.pdf")
end

if abspath(PROGRAM_FILE) == @__FILE__
  main()
end
