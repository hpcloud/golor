#!/usr/bin/env ruby
# a script to generate Go source equivalent of
# https://github.com/noprompt/vic/blob/master/lib/vic/xterm_color.rb
# NOT USED YET.

# header
puts "package xtermcolor

var SYSTEM []Color
var RGB []Color
var GREYSCALE []Color
var ALLCOLORS []Color

func init(){
    SYSTEM = []Color{
		{0x00, 0x00, 0x00}, {0x80, 0x00, 0x00}, {0x00, 0x80, 0x00},
		{0x80, 0x80, 0x00}, {0x00, 0x00, 0x80}, {0x80, 0x00, 0x80},
		{0x00, 0x80, 0x80}, {0xc0, 0xc0, 0xc0}, {0x80, 0x80, 0x80},
		{0xff, 0x00, 0x00}, {0x00, 0xff, 0x00}, {0xff, 0xff, 0x00},
		{0x00, 0x00, 0xff}, {0xff, 0x00, 0xff}, {0x00, 0xff, 0xff},
		{0xff, 0xff, 0xff},
    }
    RGB = []Color{
"

[0x00, 0x5f, 0x87, 0xaf, 0xd7, 0xff].repeated_permutation(3).to_a.each do |c|
  r,g,b = c[0].to_s(16), c[1].to_s(16), c[2].to_s(16)
  puts "        {0x#{r}, 0x#{g}, 0x#{b}},"
end

puts "    }"

puts "
    GREYSCALE = []Color {
"

(0x08..0xee).step(0x0a).map { |v| [v] * 3 }.each do |c|
  r,g,b = c[0].to_s(16), c[1].to_s(16), c[2].to_s(16)
  puts "        {0x#{r}, 0x#{g}, 0x#{b}},"
end

puts "    }
    ALLCOLORS = make([]Color, len(SYSTEM) + len(RGB) + len(GREYSCALE))
    ALLCOLORS = append(ALLCOLORS, SYSTEM...)
    ALLCOLORS = append(ALLCOLORS, RGB...)
    ALLCOLORS = append(ALLCOLORS, GREYSCALE...)
}
"
