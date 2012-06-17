using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Text;
using System.Windows.Forms;

namespace SocketIM
{
    public partial class MainForm : Form
    {
        public MainForm()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            IM im = new IM(8000, 8001);
            im.Show();
            this.button1.Enabled = false;
        }

        private void button2_Click(object sender, EventArgs e)
        {
            IM im = new IM(8001, 8000);
            im.Show();
            this.button2.Enabled = false;
        }
    }
}
