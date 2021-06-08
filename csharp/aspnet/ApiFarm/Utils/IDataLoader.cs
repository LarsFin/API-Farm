using System.Collections.Generic;

namespace ApiFarm.Utils
{
    /// <summary>
    /// Captures behaviour to load and parse data from files.
    /// </summary>
    /// <typeparam name="T">The type of data to be loaded into.</typeparam>
    public interface IDataLoader<out T>
    {
        /// <summary>
        /// Loads data of a specified type from a passed data file.
        /// </summary>
        /// <param name="dataPath">The path to the data file to load from.</param>
        /// <returns>Enumerable of implemented data type.</returns>
        IEnumerable<T> Load(string dataPath);
    }
}
