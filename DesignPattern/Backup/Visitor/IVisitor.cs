using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
namespace Visitor
{
    /// <summary>
    /// The 'Visitor' interface
    /// </summary>
    interface IVisitor
    {
        void Visit(Element element);
    }
}
