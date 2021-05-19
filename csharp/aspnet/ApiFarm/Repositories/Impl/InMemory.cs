using ApiFarm.Models;

namespace ApiFarm.Repositories
{
    /// <summary>
    /// Non permanent storage option to manage VideoGames.
    /// </summary>
    public class InMemory : IRepository<VideoGame>
    {
    }
}
