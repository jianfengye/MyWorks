using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Linq;
using System.Text;
using System.Windows.Forms;
using System.Runtime.InteropServices;

namespace MouseHookApp
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        Mouse mouse;

        private void button1_Click(object sender, EventArgs e)
        {
            mouse = new Mouse();

            mouse.MouseHookStart(OnMouseProc);
        }

        private void button2_Click(object sender, EventArgs e)
        {
            mouse.MouseHookStop();
        }

        public int OnMouseProc(int nCode, IntPtr wParam, IntPtr lParam)
        {
            int WM_LBUTTONDOWN = 0x201;
            if (wParam.ToInt32() == WM_LBUTTONDOWN)
            {
                Mouse.MouseMSG m = (Mouse.MouseMSG)Marshal.PtrToStructure(lParam, typeof(Mouse.MouseMSG));

                this.richTextBox1.Text = this.richTextBox1.Text + "mouse click\r\n";
                this.richTextBox1.Visible = true;
            }
           

            return Mouse.CallNextHookEx(Mouse.hMouseHook, nCode, wParam, lParam);
        }
    }
}
