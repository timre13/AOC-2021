package main

import (
    "io/ioutil"
    "strconv"
    "strings"
    "fmt"
    "math"
    "github.com/veandco/go-sdl2/sdl"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func readFile() string {
    bytes, err := ioutil.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }
    input := string(bytes)
    return input
}

type Line struct {
    FromX   int
    FromY   int
    ToX     int
    ToY     int
}

func (l *Line) isHorizontal() bool {
    return l.FromY == l.ToY
}

func (l *Line) isVertical() bool {
    return l.FromX == l.ToX
}

func (l *Line) fromString(str string) {
    positions := strings.Split(str, " -> ")

    fromXS := strings.Split(positions[0], ",")[0]
    fromX, err := strconv.Atoi(fromXS)
    checkErr(err)
    l.FromX = fromX

    fromYS := strings.Split(positions[0], ",")[1]
    fromY, err := strconv.Atoi(fromYS)
    checkErr(err)
    l.FromY = fromY

    toXS := strings.Split(positions[1], ",")[0]
    toX, err := strconv.Atoi(toXS)
    checkErr(err)
    l.ToX = toX

    toYS := strings.Split(positions[1], ",")[1]
    toY, err := strconv.Atoi(toYS)
    checkErr(err)
    l.ToY = toY
}

func getMaxXCoord(lines []Line) int {
    max := 0
    for _, line := range lines {
        if line.FromX > max {
            max = line.FromX
        }

        if line.ToX > max {
            max = line.ToX
        }
    }
    return max
}

func getMaxYCoord(lines []Line) int {
    max := 0
    for _, line := range lines {
        if line.FromY > max {
            max = line.FromY
        }

        if line.ToY > max {
            max = line.ToY
        }
    }
    return max
}

func countOverlapping(grid [][]int) int {
    overlappingCount := 0
    for _, row := range grid {
        for _, val := range row {
            if val >= 2 {
                overlappingCount++
            }
        }
    }
    return overlappingCount
}

func main() {
    input := readFile()

    lines := []Line{}
    for _, l := range strings.Split(input, "\n") {
        if len(l) == 0 {
            continue
        }

        line := Line{}
        line.fromString(l)
        fmt.Println(line)
        lines = append(lines, line)
    }

    maxX, maxY := getMaxXCoord(lines), getMaxYCoord(lines)
    grid := make([][]int, maxY+1)
    for i := 0; i < maxY+1; i++ {
        grid[i] = make([]int, maxX+1)
    }
    for _, line := range lines {
        if line.isHorizontal() { // Y equal
            for i := int(math.Min(float64(line.FromX), float64(line.ToX))); i <= int(math.Max(float64(line.FromX), float64(line.ToX))); i++ {
                grid[line.FromY][i]++
            }
        } else if line.isVertical() { // X equal
            for i := int(math.Min(float64(line.FromY), float64(line.ToY))); i <= int(math.Max(float64(line.FromY), float64(line.ToY))); i++ {
                grid[i][line.FromX]++
            }
        } else {
            x := line.FromX
            y := line.FromY
            x2 := line.ToX
            y2 := line.ToY
            grid[y][x]++

            for {
                if x < x2 {
                    x++
                } else if x > x2 {
                    x--
                } else {
                    break
                }

                if y < y2 {
                    y++
                } else if y > y2 {
                    y--
                } else {
                    break
                }
                grid[y][x]++
            }
        }
    }

    fmt.Printf("There are %d overlapping points\n", countOverlapping(grid))

    //---------------------------- SDL Init ------------------------------------
    err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO)
    checkErr(err)

    winW, winH := int32(maxX+1), int32(maxY+1)
    window, err := sdl.CreateWindow(
            "Advent of Code 2021 - Day 5/B",
            sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
            winW, winH,
            0)
    checkErr(err)

    renderer, err := sdl.CreateRenderer(window, 0, sdl.RENDERER_TARGETTEXTURE)
    checkErr(err)

    //-------------------------- Render into texture ---------------------------

    renderer.SetDrawColor(15, 15, 35, 255)
    renderer.Clear()
    
    tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_RGB24, sdl.TEXTUREACCESS_TARGET, winW, winH)
    checkErr(err)
    renderer.SetRenderTarget(tex)

    colors := [...][3]uint8{
        { 15,  15,  35},
        { 42,  77,  25},
        { 97, 198,  49},
        {198, 179,  49},
        {220, 190,  50},
        {255,   0,   0},
    }
    for y := 0; y < maxY+1; y++ {
        for x := 0; x < maxX+1; x++ {
            renderer.SetDrawColor(colors[grid[x][y]][0], colors[grid[x][y]][1], colors[grid[x][y]][2], 255)
            renderer.DrawPoint(int32(x), int32(y))
        }
    }

    renderer.SetRenderTarget(nil)

    //--------------------------- Main loop ------------------------------------
    scale := 1.0
    done := false
    viewXOffs, viewYOffs := int32(0), int32(0)
    isMouseBtnDown := false
    lastMouseX, lastMouseY := int32(0), int32(0)
    for {
        for {
            var event = sdl.PollEvent()
            if event == nil { // No more events in the queue
                break
            }

            switch event.GetType() {
            case sdl.QUIT:
                done = true

            case sdl.KEYDOWN:
                switch event.(*sdl.KeyboardEvent).Keysym.Sym {
                case sdl.K_1:
                    scale *= 1.1

                case sdl.K_0:
                    scale /= 1.1

                case sdl.K_UP:
                    viewYOffs -= 20

                case sdl.K_DOWN:
                    viewYOffs += 20

                case sdl.K_LEFT:
                    viewXOffs -= 20

                case sdl.K_RIGHT:
                    viewXOffs += 20

                case sdl.K_DELETE:
                    viewXOffs, viewYOffs = 0, 0
                    scale = 1.0
                }

            case sdl.MOUSEWHEEL:
                origW, origH := int32(float64(winW)*scale), int32(float64(winH)*scale)
                if event.(*sdl.MouseWheelEvent).Y > 0 {
                    scale *= 1.1
                } else {
                    scale /= 1.1
                }
                newW, newH := int32(float64(winW)*scale), int32(float64(winH)*scale)
                viewXOffs -= (origW-newW)/2
                viewYOffs -= (origH-newH)/2

            case sdl.MOUSEBUTTONDOWN:
                if event.(*sdl.MouseButtonEvent).Button == sdl.BUTTON_LEFT {
                    isMouseBtnDown = true
                }

            case sdl.MOUSEBUTTONUP:
                if event.(*sdl.MouseButtonEvent).Button == sdl.BUTTON_LEFT {
                    isMouseBtnDown = false
                }

            case sdl.MOUSEMOTION:
                if isMouseBtnDown {
                    moveX := event.(*sdl.MouseMotionEvent).X - lastMouseX
                    moveY := event.(*sdl.MouseMotionEvent).Y - lastMouseY
                    viewXOffs -= moveX
                    viewYOffs -= moveY
                }
            }
        }
        if done {
            break
        }
        lastMouseX, lastMouseY, _ = sdl.GetMouseState()

        renderer.SetDrawColor(0, 0, 0, 255)
        renderer.Clear()

        rect := sdl.Rect{
                X: -viewXOffs, Y: -viewYOffs,
                W: int32(float64(winW)*scale), H: int32(float64(winH)*scale)}
        renderer.Copy(tex, nil, &rect)

        renderer.Present()
        sdl.Delay(16)
    }

    //----------------------------- Cleanup ------------------------------------

    tex.Destroy()
    renderer.Destroy()
    window.Destroy()
    sdl.Quit()
}
