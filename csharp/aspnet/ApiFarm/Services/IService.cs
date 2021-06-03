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
        /// Requests model with identifier passed.
        /// </summary>
        /// <param name="id">The identifier of the model to retrieve.</param>
        /// <returns>The model retrieved.</returns>
        IQuery<T> Get(uint id);

        /// <summary>
        /// Requests all models from storage.
        /// </summary>
        /// <returns>Enumerable sequence of relevant models.</returns>
        IQuery<IEnumerable<T>> GetAll();

        /// <summary>
        /// Requests to add a model to storage.
        /// </summary>
        /// <param name="model">The model to add.</param>
        /// <returns>The model which was added.</returns>
        IQuery<T> Add(T model);

        /// <summary>
        /// Retrieves model with specified identifier from storage. Then updates fields using the passed model instance.
        /// The updated model is then requested to be updated in storage.
        /// </summary>
        /// <param name="id">The identifier of the model to update.</param>
        /// <param name="updateModelValues">A model which can be partially valued. The fields with values set are used to update the original instance.</param>
        /// <returns>The updated model in storage.</returns>
        IQuery<T> Update(uint id, T updateModelValues);
    }
}
