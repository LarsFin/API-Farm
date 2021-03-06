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
        /// Resets storage so primary key is zero'd and stored models are wiped.
        /// </summary>
        void Reset();

        /// <summary>
        /// Retrieves model with identifier passed.
        /// </summary>
        /// <param name="id">The identifier of the model to retrieve.</param>
        /// <returns>The desired model or its default when not in storage.</returns>
        T Get(uint id);

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

        /// <summary>
        /// Updates a model in internal storage.
        /// </summary>
        /// <param name="updatedModel">The updated model to replace the original with.</param>
        /// <returns>The updated model instance.</returns>
        T Update(T updatedModel);

        /// <summary>
        /// Deletes a model from internal storage.
        /// </summary>
        /// <param name="id">The identifier of the model to remove.</param>
        /// <returns>The model which was deleted.</returns>
        T Delete(uint id);
    }
}
