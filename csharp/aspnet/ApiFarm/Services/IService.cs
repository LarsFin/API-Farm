using System.Collections.Generic;
using ApiFarm.Models;
using ApiFarm.Utils;

namespace ApiFarm.Services
{
    /// <summary>
    /// Facilitates communication between controller and storage.
    /// </summary>
    /// <typeparam name="T">The type of the <see cref="IModel"/> which is managed via the implemented service.</typeparam>
    public interface IService<T>
        where T : IModel
    {
        /// <summary>
        /// Requests all models from storage.
        /// </summary>
        /// <returns>Enumerable sequence of relevant models.</returns>
        IQuery<IEnumerable<T>> GetAll();
    }
}
