using System.Collections.Generic;
using ApiFarm.Models.Impl;
using ApiFarm.Repositories;
using ApiFarm.Utils;

namespace ApiFarm.Services.Impl
{
    /// <summary>
    /// Interfaces storage to retrieve <see cref="VideoGame"/> elements.
    /// </summary>
    public class VideoGameService : IService<VideoGame>
    {
        private IRepository<VideoGame> videoGameStorage;
        private IQueryFactory queryFactory;

        /// <summary>
        /// Initializes a new instance of the <see cref="VideoGameService"/> class.
        /// </summary>
        /// <param name="videoGameStorage">Storage interface for managing <see cref="VideoGame"/> models.</param>
        /// <param name="queryFactory">Factory responsible for initialising Queries with <see cref="VideoGame"/> related results.</param>
        public VideoGameService(
            IRepository<VideoGame> videoGameStorage,
            IQueryFactory queryFactory)
        {
            this.videoGameStorage = videoGameStorage;
            this.queryFactory = queryFactory;
        }

        /// <summary>
        /// Retrieves all <see cref="VideoGame"/> elements from storage.
        /// </summary>
        /// <returns><see cref="VideoGame"/> elements as enumerable series.</returns>
        public IQuery<IEnumerable<VideoGame>> GetAll()
        {
            var storedVideoGames = this.videoGameStorage.GetAll();

            return this.queryFactory.Build(result: storedVideoGames);
        }

        /// <summary>
        /// Adds <see cref="VideoGame"/> to storage.
        /// </summary>
        /// <param name="videoGame">The <see cref="VideoGame"/> to add to storage.</param>
        /// <returns>The <see cref="VideoGame"/> which was added.</returns>
        public IQuery<VideoGame> Add(VideoGame videoGame)
        {
            return default;
        }
    }
}
