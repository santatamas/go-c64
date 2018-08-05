using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.InteropServices;
using System.Text;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Win32.SafeHandles;

namespace Hackaton
{
    class FastConsole
    {
        private short width_;
        private short height_;
        private SafeFileHandle handle_;
        private CharInfo[] buffer_;
        private bool[] invert_;
        private object sync_ = new object();
        private bool dirty_;

        public FastConsole(short width, short height)
        {
            Console.SetWindowSize(width, height);
            Console.SetBufferSize(width, height);

            width_ = width;
            height_ = height;
            buffer_ = new CharInfo[width*height];
            invert_ = new bool[width*height];
            for (int i = 0; i < width * height; i++)
            {
                buffer_[i] = new CharInfo() {Attributes = 7};
                invert_[i] = false;
            }

            handle_ = CreateFile("CONOUT$", 0x40000000, 2, IntPtr.Zero, FileMode.Open, 0, IntPtr.Zero);
             if (handle_.IsInvalid)
                 throw new Exception("Error opening console.");

            Thread th = new Thread(Refresh);
            th.Start();
        }

        public void SetChar(int x, int y, char ch, bool invert = false)
        {
            lock (sync_)
            {
                buffer_[x + width_ * y].Char.AsciiChar = (byte)ch;
                invert_[x + width_ * y] = invert;
                dirty_ = true;
            }
        }

        public void SetBackColor(int x, int y, ConsoleColor color)
        {
            lock (sync_)
            {
                short attribute = buffer_[x + width_ * y].Attributes;
                attribute &= 0x0F;
                buffer_[x + width_ * y].Attributes = (short)(attribute | (((int)(color)) << 4));
                dirty_ = true;
            }
        }

        public void SetForeColor(int x, int y, ConsoleColor foreColor)
        {
            lock (sync_)
            {
                short attribute = buffer_[x + width_ * y].Attributes;
                attribute &= 240;
                buffer_[x + width_ * y].Attributes = (short)((int)foreColor | attribute);
                dirty_ = true;
            }
        }

        private void Invert()
        {
            for (int i = 0; i < buffer_.Length; i++)
            {
                if (invert_[i])
                    buffer_[i].Attributes = (short)(((buffer_[i].Attributes & 0x0f) << 4) + ((buffer_[i].Attributes & 0xF0) >> 4));
            }
        }

        private void Refresh()
        {
            SmallRect rect = new SmallRect() { Left = 0, Top = 0, Right = width_, Bottom = height_ };

            for (;;)
            {
                lock(sync_)
                {
                    if (dirty_)
                    {
                        Invert();
                        WriteConsoleOutput(handle_, buffer_, new Coord() {X = width_, Y = height_},
                            new Coord() {X = 0, Y = 0}, ref rect);
                        Invert();
                        dirty_ = false;
                    }
                }
                Thread.Sleep(50);
            }
        }


        #region PInvoke

        [DllImport("Kernel32.dll", SetLastError = true, CharSet = CharSet.Auto)]
        static extern SafeFileHandle CreateFile(
            string fileName,
            [MarshalAs(UnmanagedType.U4)] uint fileAccess,
            [MarshalAs(UnmanagedType.U4)] uint fileShare,
            IntPtr securityAttributes,
            [MarshalAs(UnmanagedType.U4)] FileMode creationDisposition,
            [MarshalAs(UnmanagedType.U4)] int flags,
            IntPtr template);

        [DllImport("kernel32.dll", SetLastError = true)]
        static extern bool WriteConsoleOutput(
          SafeFileHandle hConsoleOutput,
          CharInfo[] lpBuffer,
          Coord dwBufferSize,
          Coord dwBufferCoord,
          ref SmallRect lpWriteRegion);

        [StructLayout(LayoutKind.Sequential)]
        public struct Coord
        {
            public short X;
            public short Y;

            public Coord(short X, short Y)
            {
                this.X = X;
                this.Y = Y;
            }
        };

        [StructLayout(LayoutKind.Explicit)]
        public struct CharUnion
        {
            [FieldOffset(0)]
            public char UnicodeChar;
            [FieldOffset(0)]
            public byte AsciiChar;
        }

        [StructLayout(LayoutKind.Explicit)]
        public struct CharInfo
        {
            [FieldOffset(0)]
            public CharUnion Char;
            [FieldOffset(2)]
            public short Attributes;
        }

        [StructLayout(LayoutKind.Sequential)]
        public struct SmallRect
        {
            public short Left;
            public short Top;
            public short Right;
            public short Bottom;
        }
        #endregion


        public int Width { get { return width_; } }

        public int Height { get { return height_; } }
    }
}
