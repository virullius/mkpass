require "option_parser"

module Mkpass
  VERSION = "0.2.0"
  MAX_TRIES = 10

  UPPER = "ABCDEFGHIJLKMNOPQRSTUVWXYZ"
  LOWER = "abcdefghijlkmnopqrstuvwxyz"
  NUMBER = "0123456789"
  SYMBOL = "`~!@#$%^&*()_+-=[]\\{}|;':\",./<>?"

  DefaultOptions = Options.new(true, true, true, true, 16_u16)

  struct Options
    property upper, lower, number, symbol, length

    def initialize(@upper : Bool, @lower : Bool, @number : Bool, @symbol : Bool, @length : UInt16)
    end
  end

  class Generator
    property opts, charset

    def initialize(@opts : Options)
      @charset = String.new
      self.charset += UPPER if self.opts.upper
      self.charset += LOWER if self.opts.lower
      self.charset += NUMBER if self.opts.number
      self.charset += SYMBOL if self.opts.symbol
    end

    def generate : String
      tries = 0
      str = ""
      loop do
        tries += 1
        if tries > MAX_TRIES
          raise "Failed to generate a suitable string in #{MAX_TRIES} tries"
        end

        str = ""
        i = 0
        while i < self.opts.length
          i += 1
          str += self.charset[Random.rand(self.charset.size)]
        end

        if self.opts.upper && str.each_char.map{|c| UPPER.includes?(c)}.size == 0
          next
        end
        if self.opts.lower && str.each_char.map{|c| LOWER.includes?(c)}.size == 0
          next
        end
        if self.opts.number && str.each_char.map{|c| NUMBER.includes?(c)}.size == 0
          next
        end
        if self.opts.symbol && str.each_char.map{|c| SYMBOL.includes?(c)}.size == 0
          next
        end
        break

      end
      str
    end
  end

end

opts = Mkpass::DefaultOptions
n = 0
OptionParser.parse do |p|
  p.on("-xu", "Exclude upper case characters") { opts.upper = false }
  p.on("-xl", "Exclude lower case characters") { opts.lower = false }
  p.on("-xn", "Exclude number characters") { opts.number = false }
  p.on("-xs", "Exclude symbol characters") { opts.symbol = false }
  p.on("-n Number", "Number of strings to generate") { |num| n = num.to_u16 }
  p.on("-l Length", "Length of string to generate") { |l| opts.length = l.to_u16 }
end

g = Mkpass::Generator.new(opts)
i = 0
while i < n
  puts g.generate
  i += 1
end
