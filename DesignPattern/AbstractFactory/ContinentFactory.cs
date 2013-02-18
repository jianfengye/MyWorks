using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace AbstractFactory
{
    /// <summary>
    /// AbstractFactory
    /// </summary>
    abstract class ContinentFactory
    {
        /// <summary>
        /// CreateProductA()
        /// </summary>
        public abstract Herbivor CreateHerbivore();

        /// <summary>
        /// CreateProductB()
        /// </summary>
        public abstract Carnivore CreateCarnivore();
    }

    /// <summary>
    /// AbStractProductA
    /// </summary>
    public abstract class Herbivor
    { 
    }

    /// <summary>
    /// AbStractProductB
    /// </summary>
    public abstract class Carnivore
    {
        public abstract void Eat(Herbivor h);
    }
}
