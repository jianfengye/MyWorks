using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.Text;
using System.Windows.Forms;
using System.Net;
using System.Net.Sockets;
using System.Threading;

namespace SocketIM
{
    public partial class IM : Form
    {
        private int listenPort;
        private int sendPort;

        delegate void SetListBox(string strValue);                   //定义委托

        public IM()
        {
            InitializeComponent();
        }

        public IM(int listenPort, int sendPort)
        {
            InitializeComponent();
            this.listenPort = listenPort;
            this.sendPort = sendPort;
            //BeginListen();
            Thread thread = new Thread(new ThreadStart(BeginListen));
            thread.IsBackground = true;
            thread.Start();
        }

        private void BeginListen()
        {
            IPEndPoint endpoint = new IPEndPoint(IPAddress.Parse("127.0.0.1"), this.listenPort);
            Socket socket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
            byte[] byteMessage = new byte[100];

            socket.Bind(endpoint);
            SetListBoxValue("This client is listening ...");
            socket.Listen(5);

            while (true)
            {
                Socket clientSocket = socket.Accept();
                clientSocket.Receive(byteMessage);

                string messageHead = DateTime.Now.ToShortTimeString() + " Message From " + sendPort.ToString() + ": ";
                string message = messageHead + Encoding.Default.GetString(byteMessage);

                SetListBoxValue(message);
            }
        }

        private void SetListBoxValue(string value)
        {
            if (this.listBox1.InvokeRequired)
            {
                SetListBox sbox = new SetListBox(SetListBoxValue);
                listBox1.Invoke(sbox, value);
            } 
            else 
            {
                    listBox1.Items.Add(value);
            }
        }

        private void button1_Click(object sender, EventArgs e)
        {
            string message = this.richTextBox2.Text;
            IPEndPoint endpoint = new IPEndPoint(IPAddress.Parse("127.0.0.1"), this.sendPort);
            byte[] byteMessage = Encoding.ASCII.GetBytes(message);
            Socket socket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
            socket.Connect(endpoint);
            socket.Send(byteMessage);
            socket.Shutdown(SocketShutdown.Both);
            socket.Close();

            this.richTextBox2.Text = "";
        }
    }
}
