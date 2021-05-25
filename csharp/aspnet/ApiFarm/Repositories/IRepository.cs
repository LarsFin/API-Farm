using System.Collections.Generic;
using ApiFarm.Models;

namespace ApiFarm.Repositories
{
    /// <summary>
    /// Captures required behaviours of a storage facilitator.
    /// </summary>
    /// <typeparam name="T">Type of stored element.</typeparam>
    public interface IRepository<out T>
        where T : IModel, new()
    {
        /// <summary>
        /// Retrieves enumerator containing all stored elements.
        /// </summary>
        /// <returns>Stored elements for enumeration.</returns>
        IEnumerable<T> GetAll();
    }
}
