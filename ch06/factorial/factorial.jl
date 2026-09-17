function factorial(n::Int64)
  if n < 2
    for frame in stacktrace()
      println(frame)
    end
    return 1
  else
    return n * factorial(n-1)
  end
end

function main()
  println(factorial(5))
end

main()
