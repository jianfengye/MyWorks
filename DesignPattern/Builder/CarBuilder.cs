using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Builder
{
    class CarBuilder:VehicleBuilder
    {
        public CarBuilder()
        {
            vehicle = new Vehicle("Car");
        }

        public override void BuildFrame()
        {
            vehicle["frame"] = "Car Frame";
        }

        public override void BuildEngine()
        {
            vehicle["engine"] = "2500 CC";
        }

        public override void BuildWheel()
        {
            vehicle["wheel"] = "4";
        }

        public override void BuildDoors()
        {
            vehicle["doors"] = "4";
        }
    }
}
