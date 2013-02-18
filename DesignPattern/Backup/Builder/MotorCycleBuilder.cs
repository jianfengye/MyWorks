using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Builder
{
    class MotorCycleBuilder:VehicleBuilder
    {
        public MotorCycleBuilder()
        {
            vehicle = new Vehicle("MotorCycle");
        }

        public override void BuildFrame()
        {
            vehicle["frame"] = "MotorCycle Frame";
        }

        public override void BuildEngine()
        {
            vehicle["engine"] = "MotorCycle Engine";
        }

        public override void BuildWheel()
        {
            vehicle["wheel"] = "2";
        }

        public override void BuildDoors()
        {
            vehicle["doors"] = "0";
        }
    }
}
