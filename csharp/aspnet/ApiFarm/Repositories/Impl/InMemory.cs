﻿using System.Collections.Generic;
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
        private ICloner<T> cloner;
        private List<T> videoGames;

        /// <summary>
        /// Initializes a new instance of the <see cref="InMemory{T}"/> class.
        /// </summary>
        /// <param name="cloner">Responsible for cloning stored entities to avoid external storage manipulation.</param>
        public InMemory(ICloner<T> cloner)
        {
            this.cloner = cloner;
            this.videoGames = new List<T>();
        }

        internal InMemory(ICloner<T> cloner, List<T> presetVideoGames)
        {
            this.cloner = cloner;
            this.videoGames = presetVideoGames;
        }

        /// <summary>
        /// Returns private list of stored entities.
        /// </summary>
        /// <returns>List of stored entities.</returns>
        public IEnumerable<T> GetAll()
        {
            return this.videoGames.Select(q => this.cloner.Clone(q));
        }
    }
}
