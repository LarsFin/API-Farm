namespace ApiFarm.Utils
{
    /// <summary>
    /// Signs the behaviour of a Factory instantiating <see cref="IQuery{T}"/> elements.
    /// </summary>
    public interface IQueryFactory
    {
        /// <summary>
        /// Behaviour to instantiate an instance of an implemented <see cref="IQuery{T}"/>.
        /// </summary>
        /// <typeparam name="T">The type of the entity queried.</typeparam>
        /// <param name="code">The status of the <see cref="IQuery{T}"/> made (e.g; 0: successful, 400: bad request).</param>
        /// <param name="message">The message of the <see cref="IQuery{T}"/> (e.g; 'Deleted entity with id x').</param>
        /// <param name="result">The entity which was queried.</param>
        /// <returns>The <see cref="IQuery{T}"/> implementation which was constructed.</returns>
        IQuery<T> Build<T>(uint code = default, string message = default, T result = default);
    }
}
