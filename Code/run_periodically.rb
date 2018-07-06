  def every_n_seconds(n)
    loop do
      before = Time.now
      yield
      interval = n-(Time.now-before)
      sleep(interval) if interval > 0
    end
  end

  every_n_seconds(5) do
   puts "#{Time.now.strftime("%X")}"
  end

