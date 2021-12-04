#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <cstring>
#include <array>
#include <vector>
#include <iomanip>

#define COLOR_GREEN "\033[32m"
#define COLOR_RESET "\033[0m"

inline std::stringstream readInputFile()
{
    std::fstream file;
    file.open("../input.txt");
    if (file.fail())
    {
        throw std::runtime_error{std::string("Failed to open input file: ")+std::strerror(errno)};
    }
    std::stringstream ss;
    ss << file.rdbuf();
    return ss;
}

inline std::vector<int> parseNumbers(const std::string& str)
{
    std::vector<int> numbers;
    for (size_t i{}; i < str.size();)
    {
        size_t pos;
        numbers.push_back(std::stoi(str.substr(i), &pos));
        i += pos+1;
    }
    return numbers;
}

class Board
{
private:
    using row_t = std::array<int, 5>;
    using board_t = std::array<row_t, 5>;
    board_t m_board;

    using mRow_t = std::array<bool, 5>;
    using mBoard_t = std::array<mRow_t, 5>;
    mBoard_t m_mBoard;

public:
    Board()
    {
        for (int i{}; i < 5; ++i)
        {
            for (int j{}; j < 5; ++j)
            {
                m_board[i][j] = -1;
            }
        }

        for (int i{}; i < 5; ++i)
        {
            for (int j{}; j < 5; ++j)
            {
                m_mBoard[i][j] = false;
            }
        }
    }

    void parseRow(int rowI, const std::string& row)
    {
        int colI{};
        for (size_t i{}; i < row.size(); ++colI)
        {
            size_t pos;
            m_board[rowI][colI] = std::stoi(row.substr(i), &pos);
            i += pos+1;
        }
    }

    void markNum(int num)
    {
        for (int i{}; i < 5; ++i)
        {
            for (int j{}; j < 5; ++j)
            {
                if (m_board[i][j] == num)
                {
                    m_mBoard[i][j] = true;
                }
            }
        }
    }

    bool hasWon()
    {
        auto isRowComplete{[&](int row){ // -> bool
            for (int i{}; i < 5; ++i)
            {
                if (!m_mBoard[row][i])
                {
                    return false;
                }
            }
            return true;
        }};

        auto isColComplete{[&](int col){ // -> bool
            for (int i{}; i < 5; ++i)
            {
                if (!m_mBoard[i][col])
                {
                    return false;
                }
            }
            return true;
        }};

        for (int i{}; i < 5; ++i)
        {
            if (isRowComplete(i) || isColComplete(i))
            {
                return true;
            }
        }
        return false;
    }

    int getUnmarkedSum()
    {
        int sum{};
        for (int i{}; i < 5; ++i)
        {
            for (int j{}; j < 5; ++j)
            {
                if (!m_mBoard[i][j])
                    sum += m_board[i][j];
            }
        }
        return sum;
    }

    friend std::ostream& operator<<(std::ostream& out, const Board& b)
    {
        for (int i{}; i < 5; ++i)
        {
            for (int j{}; j < 5; ++j)
            {
                int val = b.m_board[i][j];
                if (b.m_mBoard[i][j])
                    out << COLOR_GREEN;
                else
                    out << COLOR_RESET;
                out << std::setw(2) << std::setfill(' ') << val;
                if (j != 4)
                    out << ' ';
            }
            out << '\n';
        }
        out << COLOR_RESET;
        return out;
    }
};

int main()
{
    std::stringstream input;
    try
    {
        input = readInputFile();
    }
    catch (std::exception& e)
    {
        std::cerr << e.what() << '\n';
        return 1;
    }

    std::string numbersStr;
    std::getline(input, numbersStr);
    auto numbers = parseNumbers(numbersStr);


    std::vector<Board> boards;

    std::string line;
    std::getline(input, line);
    while (true)
    {
        Board b;
        for (int i{}; i < 5; ++i)
        {
            std::getline(input, line);
            b.parseRow(i, line);
        }
        //std::cout << b << '\n';
        boards.push_back(std::move(b));

        if (!std::getline(input, line))
            break;
    }

    for (int num : numbers)
    {
        std::cout << "\n\nMarking: " << num << '\n';
        for (size_t i{}; i < boards.size(); ++i)
        {
            Board& b = boards[i];
            b.markNum(num);
            std::cout << "Board #" << (i+1) << '\n' << b << '\n';
            if (b.hasWon())
            {
                std::cout << "Score: " << (b.getUnmarkedSum()*num) << '\n';
                return 0;
            }
        }
    }

    return 0;
}
