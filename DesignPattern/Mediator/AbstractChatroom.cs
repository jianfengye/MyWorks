using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
namespace Mediator
{

    /// <summary>
    /// The 'Mediator' abstract class
    /// </summary>
    abstract class AbstractChatroom
    {
        public abstract void Register(Participant participant);
        public abstract void Send(
          string from, string to, string message);
    }
}
