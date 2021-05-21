namespace ApiFarm.Utils
{
    /// <summary>
    /// Signs the available properties of a query made against a service and storage.
    /// </summary>
    /// <typeparam name="T">The type of the element being queried.</typeparam>
    public interface IQuery<out T>
    {
        /// <summary>
        /// Gets the status of the query (e.g; 200: ok, 404: not found).
        /// </summary>
        uint Code { get; }

        /// <summary>
        /// Gets the message of the query (e.g; 'Entity with id x was deleted.').
        /// </summary>
        string Message { get; }

        /// <summary>
        /// Gets the entity or entities from a query.
        /// </summary>
        T Result { get; }
    }
}
