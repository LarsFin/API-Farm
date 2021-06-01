using System.Collections.Generic;

namespace ApiFarm.Utils
{
    /// <summary>
    /// Includes extension methods to neaten logic.
    /// </summary>
    public static class Extensions
    {
        /// <summary>
        /// Wraps equality comparer for clearer syntax option.
        /// </summary>
        /// <typeparam name="T">Type of object to be checked.</typeparam>
        /// <param name="obj">The object to check whether is default.</param>
        /// <returns>Whether the checked object is the default value of its type.</returns>
        public static bool IsDefault<T>(this T obj) => EqualityComparer<T>.Default.Equals(obj, default);
    }
}
