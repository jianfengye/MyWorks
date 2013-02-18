using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Chain_of_Responsibility
{
    class Program
    {
        /// <summary>

        /// Entry point into console application.

        /// </summary>

        static void Main()
        {

            // Setup Chain of Responsibility

            Approver larry = new Director();

            Approver sam = new VicePresident();

            Approver tammy = new President();



            larry.SetSuccessor(sam);

            sam.SetSuccessor(tammy);



            // Generate and process purchase requests

            Purchase p = new Purchase(2034, 350.00, "Assets");

            larry.ProcessRequest(p);



            p = new Purchase(2035, 32590.10, "Project X");

            larry.ProcessRequest(p);



            p = new Purchase(2036, 122100.00, "Project Y");

            larry.ProcessRequest(p);



            // Wait for user

            Console.ReadKey();

        }

    }
}
