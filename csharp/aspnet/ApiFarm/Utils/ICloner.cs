using ApiFarm.Models;

namespace ApiFarm.Utils
{
    /// <summary>
    /// Includes method to deep clone instance of type defined by implementation.
    /// </summary>
    /// <typeparam name="T">The type of the entity to be cloned.</typeparam>
    public interface ICloner<T>
        where T : IModel
    {
        /// <summary>
        /// Creates a deep copy of an original instance.
        /// </summary>
        /// <param name="original">The original instance to clone.</param>
        /// <returns>The deep clone which should have new references.</returns>
        T Clone(T original);
    }
}
