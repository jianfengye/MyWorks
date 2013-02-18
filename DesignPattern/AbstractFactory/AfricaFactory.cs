using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace AbstractFactory
{
    /// <summary>
    /// ConcreteFactory1
    /// </summary>
    class AfricaFactory:ContinentFactory
    {
        public override Herbivor CreateHerbivore()
        {
            return new Wildbeest();
        }

        public override Carnivore CreateCarnivore()
        {
            return new Lion();
        }
    }

    /// <summary>
    /// ProductA1
    /// </summary>
    class Wildbeest : Herbivor
    { 
    }

    /// <summary>
    /// ProductB1
    /// </summary>
    class Lion : Carnivore
    {
        public override void Eat(Herbivor h)
        {
            Console.WriteLine(this.GetType().Name + " eats " + h.GetType().Name);
        }
    }
}
