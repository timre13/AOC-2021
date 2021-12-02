local sdl = require "SDL"

local file = io.open("input.txt", "r")
if file == nil then error("Failed to open input file\n") end
local input = file:read("a")

local hpos = 0
local depth = 0
local aim = 0

local depthVals = {}
local hposVals = {}

for line in input:gmatch("[^\r\n]+") do
    --print(line)
    local cmd = line:sub(0, line:find(" ")-1)
    local arg = tonumber(line:sub(line:find(" ")+1, -1))
    if arg == nil then error("Failed to convert argument to number") end

    if cmd == "forward" then
        hpos = hpos + arg
        depth = depth + aim * arg
    elseif cmd == "down" then
        aim = aim + arg
    elseif cmd == "up" then
        aim = aim - arg
    end

    table.insert(depthVals, depth)
    table.insert(hposVals, hpos)
end

print("Depth:           "..depth)
print("Horizontal pos.: "..hpos)
print("Product:         "..(depth*hpos))

local ret, err = sdl.init {
    sdl.flags.Video
}
if not ret then error(err) end

local win, err = sdl.createWindow {
    title = "AoC 2021 - Day 2b Visualization - depth: "..depth..", hpos: "..hpos..", prod.: "..(depth*hpos),
    width = 1500,
    height = 800
}
if not win then error(err) end

local rend, err = sdl.createRenderer(win, 0, 0)
if not rend then error(err) end

local function max(t)
    local m = 0
    for _, val in ipairs(t) do
        if val > m then
            m = val
        end
    end
    return m
end

local maxHpos = max(hposVals)
local maxDepth = max(depthVals)

local running = true
while running do
    for e in sdl.pollEvent() do
        if e.type == sdl.event.Quit then
            running = false
        end
    end

    rend:setDrawColor{r=15, g=15, b=35}
    rend:clear()

    local prevDepth = 0
    local prevHpos = 0
    for i, d in ipairs(depthVals) do
        local h = hposVals[i]

        rend:setDrawColor{r=0, g=153, b=0}
        rend:drawLine{
            x1=prevHpos/maxHpos*1500//1, y1=prevDepth/maxDepth*800//1,
            x2=h/maxHpos*1500//1, y2=d/maxDepth*800//1}

        prevDepth = d
        prevHpos = h
    end

    rend:present()
    sdl.delay(16)
end

sdl.quit()
