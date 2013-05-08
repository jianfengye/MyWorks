using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Runtime.InteropServices;
using System.Reflection;

namespace MouseHookApp
{
    class Mouse
    {
        public delegate int MouseProc(int nCode, IntPtr wParam, IntPtr lParam);

        MouseProc MouseHookProcedure;

        public static int hMouseHook = 0;
        public const int WH_MOUSE_LL = 14;

        public struct Point
        {
            public long x;
            public long y;
        }

        public struct MouseMSG
        {
            public Point pt;
            public IntPtr hwnd;
            public uint wHitTestCode;
            public IntPtr dwExtraInfo;
        }

        [DllImport("user32.dll", CharSet = CharSet.Auto, CallingConvention = CallingConvention.StdCall)]
        public static extern int SetWindowsHookEx(int idHook, MouseProc lpfn, IntPtr hInstance, int threadId);

        [DllImport("user32.dll", CharSet = CharSet.Auto, CallingConvention = CallingConvention.StdCall)]
        public static extern bool UnhookWindowsHookEx(int idHook);

        [DllImport("user32.dll", CharSet = CharSet.Auto, CallingConvention = CallingConvention.StdCall)]
        public static extern int CallNextHookEx(int idHook, int nCode, IntPtr wParam, IntPtr lParam);

        public void MouseHookStart(MouseProc onMouseProc)
        {
            if (hMouseHook == 0)
            {
                MouseHookProcedure = new MouseProc(onMouseProc);

                // 设置钩子
                hMouseHook = SetWindowsHookEx(WH_MOUSE_LL, MouseHookProcedure, IntPtr.Zero, 0);

                if (hMouseHook == 0)
                {
                    MouseHookStop();
                    throw new Exception("SetWindowsHookEx failed");
                }
            }
        }

        public void MouseHookStop()
        {
            bool unHookRet = true;
            if (hMouseHook != 0)
            {
                unHookRet = UnhookWindowsHookEx(hMouseHook);
                hMouseHook = 0;
            }

            // 异常处理
            if (!(unHookRet))
            {
                throw new Exception("UnhookWindwosHookEx failed");
            }
        }
    }
}
