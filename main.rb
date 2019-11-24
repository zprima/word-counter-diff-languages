file_path = './moby.txt'

words = 0
in_word = false

File.open(file_path) do |file|
  until file.eof?

    buffer = file.read(1)
    
    if buffer == ' ' && in_word
      words += 1
      in_word = false
    end

    in_word = buffer.match?(/\A[a-zA-Z'-]*\z/)
  end
end

puts "Words counted: #{words}"