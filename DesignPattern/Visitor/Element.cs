using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
namespace Visitor
{
    /// <summary>
    /// The 'Element' abstract class
    /// </summary>
    abstract class Element
    {
        public abstract void Accept(IVisitor visitor);
    }
}
