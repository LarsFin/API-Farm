using ApiFarm.Models.Impl;

namespace ApiFarm.Utils.Impl
{
    /// <summary>
    /// Responsible for cloning <see cref="VideoGame"/> instances.
    /// </summary>
    public class VideoGameCloner : ICloner<VideoGame>
    {
        /// <summary>
        /// Forms a deep clone of a <see cref="VideoGame"/>.
        /// </summary>
        /// <param name="original">The original <see cref="VideoGame"/> to clone.</param>
        /// <returns>The copied <see cref="VideoGame"/>.</returns>
        public VideoGame Clone(VideoGame original)
        {
            return default;
        }
    }
}
