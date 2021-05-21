using System.Collections.Generic;
using System.Runtime.CompilerServices;
using ApiFarm.Models;

[assembly: InternalsVisibleTo("ApiFarm.Tests")]

namespace ApiFarm.Repositories
{
    /// <summary>
    /// Non permanent storage option to manage entity.
    /// </summary>
    /// <typeparam name="T">Type of in memory stored elemnt.</typeparam>
    public class InMemory<T> : IRepository<T>
         where T : IModel, new()
    {
        private List<T> videoGames;

        /// <summary>
        /// Initializes a new instance of the <see cref="InMemory{T}"/> class.
        /// </summary>
        public InMemory()
        {
            this.videoGames = new List<T>();
        }

        internal InMemory(List<T> presetVideoGames)
        {
            this.videoGames = presetVideoGames;
        }

        /// <summary>
        /// Returns private list of stored entities.
        /// </summary>
        /// <returns>List of stored entities.</returns>
        public IEnumerable<T> GetAll() => this.videoGames;
    }
}
