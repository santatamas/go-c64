using System.Collections.Generic;

namespace Hackaton
{
    /// <summary>
    /// Converts C64 screen character codes to ASCII.
    /// </summary>
    public static class C64CharConverter
    {
        private const string codes = @"
      @| |  0
      A|a|  1
      B|b|  2
      I|i|  9
      J|j| 10
      K|k| 11
      L|l| 12
      M|m| 13
      N|n| 14
      O|o| 15
      P|p| 16
      Q|q| 17
      R|r| 18
      S|s| 19
      T|t| 20
      U|u| 21
      V|v| 22
      W|w| 23
      X|x| 24
      Y|y| 25
      Z|z| 26
      [| | 27
       | | 28
      ]| | 29
      ^| | 30
       | | 31  
  SPACE| | 32
      !| | 33
     ""| | 34
      #| | 35
      $| | 36
       | | 93
       | | 94
       | | 95
  SPACE| | 96
       | | 97
       | | 98
       | | 99
       | |100
       | |101
       | |102
       | |103
       | |104
      C|c|  3   
      D|d|  4   
      E|e|  5   
      %| | 37   
      &| | 38   
      '| | 39   
      (| | 40   
      )| | 41   
      *| | 42   
      +| | 43   
      ,| | 44   
      -| | 45   
      .| | 46   
      /| | 47   
      0| | 48   
      1| | 49   
      2| | 50   
      3| | 51   
      4| | 52   
      5| | 53   
      6| | 54   
      7| | 55   
      8| | 56   
      9| | 57   
      :| | 58   
      ;| | 59   
      <| | 60   
      =| | 61   
      >| | 62   
      ?| | 63   
       | | 64   
       | |105   
       | |106   
       | |107   
       | |108   
       | |109   
       | |110   
       | |111   
       | |112   
       | |113   
       | |114   
       | |115   
       | |116   
      F|f|  6
      G|g|  7
      H|h|  8
       |A| 65
       |B| 66
       |C| 67
       |D| 68
       |E| 69
       |F| 70
       |G| 71
       |H| 72
       |I| 73
       |J| 74
       |K| 75
       |L| 76
       |M| 77
       |N| 78
       |O| 79
       |P| 80
       |Q| 81
       |R| 82
       |S| 83
       |T| 84
       |U| 85
       |V| 86
       |W| 87
       |X| 88
       |Y| 89
       |Z| 90
       | | 91
       | | 92
       | |117
       | |118
       | |119
       | |120
       | |121
       | |122
       | |123
       | |124
       | |125
       | |126
       | |127";

        private static Dictionary<byte, char> byteToChar;

        private static void ParseString()
        {
            byteToChar = new Dictionary<byte, char>();
            string[] lines = codes.Split('\n');
            foreach (var line in lines)
            {
                string[] parts = line.Split('|');
                if (parts.Length<3)
                    continue;
                string p1 = parts[0].Trim().Replace("SPACE", " ");
                string p3 = parts[2].Trim();

                byteToChar.Add(byte.Parse(p3), string.IsNullOrEmpty(p1) ? '*' : p1[0]);
            }
        }

        public static char ConvertToAscii(byte value)
        {
            if (byteToChar==null)
                ParseString();
            value = (byte)(value & 127);
            return byteToChar[value];
        }
    }
}