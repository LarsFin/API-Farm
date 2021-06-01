using System.Collections.Generic;
using ApiFarm.Models;

namespace ApiFarm.Repositories
{
    /// <summary>
    /// Captures required behaviours of a storage facilitator.
    /// </summary>
    /// <typeparam name="T">Type of stored element.</typeparam>
    public interface IRepository<T>
        where T : IModel
    {
        /// <summary>
        /// Retrieves enumerator containing all stored elements.
        /// </summary>
        /// <returns>Stored elements for enumeration.</returns>
        IEnumerable<T> GetAll();

        /// <summary>
        /// Adds a model to the internal storage.
        /// </summary>
        /// <param name="model">The model to be added.</param>
        /// <returns>The model added with an amended identifier value.</returns>
        T Add(T model);
    }
}
