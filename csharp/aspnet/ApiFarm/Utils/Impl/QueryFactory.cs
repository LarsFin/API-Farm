namespace ApiFarm.Utils.Impl
{
    /// <summary>
    /// Instantiates instances of a <see cref="IQuery{T}"/> with a later determined result type.
    /// </summary>
    public class QueryFactory : IQueryFactory
    {
        /// <summary>
        /// Instantiates a <see cref="Query{T}"/> where its result type is determined on call.
        /// </summary>
        /// <typeparam name="T">The type of the entity being queried.</typeparam>
        /// <param name="code">The status of the query (e.g; 0: successful, 404: not found).</param>
        /// <param name="message">The message attached to the query (e.g; 'Could not find entity with id x').</param>
        /// <param name="result">The entity returned from the query made.</param>
        /// <returns>The constructed query.</returns>
        public IQuery<T> Build<T>(uint code = 0, string message = default(string), T result = default(T))
        {
            return new Query<T>(code, message, result);
        }
    }
}
