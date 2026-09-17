function fibonaci(n::Int)
  if (n < 3)
    return 1
  else
    return fibonaci(n - 1) + fibonaci(n - 2)
  end
end

function main()
  println(fibonaci(5))
end

main()
