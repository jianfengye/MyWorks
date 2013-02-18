using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace AbstractFactory
{
    /// <summary>
    /// ConcreteFactory2
    /// </summary>
    class AmericaFactory:ContinentFactory
    {
        public override Herbivor CreateHerbivore()
        {
            return new Biston();

        }

        public override Carnivore CreateCarnivore()
        {
            return new Wolf();
        }
    }

    /// <summary>
    /// ProductA2
    /// </summary>
    public class Biston : Herbivor
    { 
    }

    /// <summary>
    /// ProductB2
    /// </summary>
    public class Wolf : Carnivore
    {
        public override void Eat(Herbivor h)
        {
            Console.WriteLine(this.GetType().Name + " eats " + h.GetType().Name);
        }
    }
}
