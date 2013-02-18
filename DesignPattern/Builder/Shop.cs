using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace Builder
{
    class Shop
    {
        public void Construct(VehicleBuilder vehicleBuilder)
        {
            vehicleBuilder.BuildFrame();
            vehicleBuilder.BuildEngine();
            vehicleBuilder.BuildWheel();
            vehicleBuilder.BuildDoors();
        }
    }
}
