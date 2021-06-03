﻿using System.Collections.Generic;
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
        /// Retrieves <see cref="VideoGame"/> with passed identifier.
        /// </summary>
        /// <param name="id">The identifier of the <see cref="VideoGame"/>.</param>
        /// <returns><see cref="VideoGame"/> element with queried identifier.</returns>
        public IQuery<VideoGame> Get(uint id)
        {
            var storedVideoGame = this.videoGameStorage.Get(id);

            if (storedVideoGame is null)
            {
                return this.queryFactory.Build<VideoGame>(404, ResponseMessages.VideoGame.NotFound(id));
            }

            return this.queryFactory.Build(result: storedVideoGame);
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
            if (string.IsNullOrEmpty(videoGame.Name))
            {
                return this.queryFactory.Build<VideoGame>(400, ResponseMessages.VideoGame.RequiresName);
            }

            if (videoGame.DateReleased.IsDefault())
            {
                return this.queryFactory.Build<VideoGame>(400, ResponseMessages.VideoGame.RequiresDateReleased);
            }

            var storedVideoGame = this.videoGameStorage.Add(videoGame);
            return this.queryFactory.Build(result: storedVideoGame);
        }

        /// <summary>
        /// Updates <see cref="VideoGame"/> from storage with set fields on passed instance. The updated <see cref="VideoGame"/> is
        /// requested to be updated in storage.
        /// </summary>
        /// <param name="id">The identifier of the <see cref="VideoGame"/> to update.</param>
        /// <param name="updateVideoGameValues">The <see cref="VideoGame"/> with fields to update the original instance with.</param>
        /// <returns>Updated <see cref="VideoGame"/> instance.</returns>
        public IQuery<VideoGame> Update(uint id, VideoGame updateVideoGameValues)
        {
            return default;
        }
    }
}
