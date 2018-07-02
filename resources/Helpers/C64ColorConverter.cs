using System;

namespace Hackaton
{
    /// <summary>
    /// Converts C64 color values to ConsoleColor.
    /// </summary>
    public static class C64ColorConverter
    {
        public static ConsoleColor ByteToColor(byte value)
        {
            value = (byte)(value & 15);
            switch (value)
            {
                case 0:
                    return ConsoleColor.Black;
                case 1:
                    return ConsoleColor.White;
                case 2:
                    return ConsoleColor.DarkRed;
                case 3:
                    return ConsoleColor.Cyan;
                case 4:
                    return ConsoleColor.Magenta;
                case 5:
                    return ConsoleColor.DarkGreen;
                case 6:
                    return ConsoleColor.DarkBlue;
                case 7:
                    return ConsoleColor.Yellow;
                case 8:
                    return ConsoleColor.DarkYellow;
                case 9:
                    return ConsoleColor.DarkMagenta;
                case 10:
                    return ConsoleColor.Red;
                case 11:
                    return ConsoleColor.DarkGray;
                case 12:
                    return ConsoleColor.Gray;
                case 13:
                    return ConsoleColor.Green;
                case 14:
                    return ConsoleColor.Blue;
                case 15:
                    return ConsoleColor.DarkCyan;


            }
            throw new ArgumentOutOfRangeException();
        }

    }
}