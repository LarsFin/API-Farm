using System.Collections.Generic;
using System.Linq;
using System.Runtime.CompilerServices;
using ApiFarm.Models;
using ApiFarm.Utils;

[assembly: InternalsVisibleTo("ApiFarm.Test")]

namespace ApiFarm.Repositories
{
    /// <summary>
    /// Non permanent storage option to manage entity.
    /// </summary>
    /// <typeparam name="T">Type of in memory stored elemnt.</typeparam>
    public class InMemory<T> : IRepository<T>
         where T : IModel
    {
        private uint id = 0;
        private ICloner<T> cloner;
        private List<T> models;

        /// <summary>
        /// Initializes a new instance of the <see cref="InMemory{T}"/> class.
        /// </summary>
        /// <param name="cloner">Responsible for cloning stored entities to avoid external storage manipulation.</param>
        public InMemory(ICloner<T> cloner)
        {
            this.cloner = cloner;
            this.models = new List<T>();
        }

        internal InMemory(ICloner<T> cloner, List<T> presetModels)
        {
            this.cloner = cloner;
            this.models = presetModels;
        }

        /// <summary>
        /// Clears list of models and resets id counter.
        /// </summary>
        public void Reset()
        {
            this.id = 0;
            this.models.Clear();
        }

        /// <summary>
        /// Returns clone of model with passed Id.
        /// </summary>
        /// <param name="id">The identifier of the model to retrieve.</param>
        /// <returns>A clone of the desired model.</returns>
        public T Get(uint id)
        {
            var model = this.models.FirstOrDefault(q => q.Id == id);

            if (model.IsDefault())
            {
                return default;
            }

            return this.cloner.Clone(model);
        }

        /// <summary>
        /// Returns private list of stored entities.
        /// </summary>
        /// <returns>List of stored entities.</returns>
        public IEnumerable<T> GetAll()
        {
            return this.models.Select(q => this.cloner.Clone(q));
        }

        /// <summary>
        /// Adds a clone of the model to private list.
        /// </summary>
        /// <param name="model">The model to add.</param>
        /// <returns>A clone of the model to avoid mutating stored instances.</returns>
        public T Add(T model)
        {
            model.Id = ++this.id;
            var modelToStore = this.cloner.Clone(model);
            this.models.Add(modelToStore);
            return model;
        }

        /// <summary>
        /// Replaces the model in list with a cloned updated instance with the same identifier.
        /// </summary>
        /// <param name="updatedModel">The model to replace the original with.</param>
        /// <returns>A clone of the updated entity.</returns>
        public T Update(T updatedModel)
        {
            for (var i = 0; i < this.models.Count; i++)
            {
                if (this.models[i].Id == updatedModel.Id)
                {
                    this.models[i] = this.cloner.Clone(updatedModel);
                    return updatedModel;
                }
            }

            return default;
        }
    }
}
