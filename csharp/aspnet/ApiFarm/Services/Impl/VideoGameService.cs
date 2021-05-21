using System.Collections.Generic;
using ApiFarm.Models.Impl;
using ApiFarm.Utils;

namespace ApiFarm.Services.Impl
{
    /// <summary>
    /// Interfaces storage to retrieve <see cref="VideoGame"/> elements.
    /// </summary>
    public class VideoGameService : IService<VideoGame>
    {
        private IQueryFactory queryFactory;

        /// <summary>
        /// Initializes a new instance of the <see cref="VideoGameService"/> class.
        /// </summary>
        /// <param name="queryFactory">Factory responsible for initialising Queries with <see cref="VideoGame"/> related results.</param>
        public VideoGameService(IQueryFactory queryFactory)
        {
            this.queryFactory = queryFactory;
        }

        /// <summary>
        /// Retrieves all <see cref="VideoGame"/> elements from storage.
        /// </summary>
        /// <returns><see cref="VideoGame"/> elements as enumerable series.</returns>
        public IQuery<IEnumerable<VideoGame>> GetAll()
        {
            return default;
        }
    }
}
