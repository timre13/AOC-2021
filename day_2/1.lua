local file = io.open("input.txt", "r")
if file == nil then error("Failed to open input file\n") end
local input = file:read("a")

local hpos = 0
local depth = 0

for line in input:gmatch("[^\r\n]+") do
    --print(line)
    local cmd = line:sub(0, line:find(" ")-1)
    local arg = tonumber(line:sub(line:find(" ")+1, -1))
    if arg == nil then error("Failed to convert argument to number") end

    if cmd == "forward" then
        hpos = hpos + arg
    elseif cmd == "down" then
        depth = depth + arg
    elseif cmd == "up" then
        depth = depth - arg
    end
end

print("Depth:           "..depth)
print("Horizontal pos.: "..hpos)
print("Product:         "..(depth*hpos))
