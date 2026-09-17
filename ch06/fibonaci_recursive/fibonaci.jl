function fibonacci(n::Int)
  a = 0
  b = 1

  while n > 0
    a, b = b, a + b
    n -= 1
  end

  return a
end

println(fibonacci(5))
